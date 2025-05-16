#[cfg(not(feature = "library"))]
use cosmwasm_std::{
    Deps, DepsMut, Env, MessageInfo, Response, entry_point,
};
use log::info;
use cw2::set_contract_version;
use crate::error::PoolError;
use crate::msg::{ExecuteMsg, QueryBlockEmissionResponse, InstantiateMsg, QueryMsg};
use crate::state::{State, STATE};
use qcore_bindings::{QcoreMsg, QcoreQuerier, QcoreQuery};
use cosmwasm_std::{
    to_json_binary, Binary, StdResult,
};

// version info for migration info
const CONTRACT_NAME: &str = "crates.io:pool-querier";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut<QcoreQuery>,
    _env: Env,
    info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, PoolError> {
    let state = State {
        owner: info.sender.clone(),
    };
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    STATE.save(deps.storage, &state)?;

    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("owner", info.sender))
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut<QcoreQuery>,
    _env: Env,
    _info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response<QcoreMsg>, PoolError> {
    match msg {
        ExecuteMsg::MsgMintTribute {
            creator,
            contract_address,
            mint_amount,
            receipt_address,
        } => msg_mint_tribute(deps, creator, contract_address, mint_amount, receipt_address),
    }
}

pub fn msg_mint_tribute(
    _: DepsMut<QcoreQuery>,
    creator: String,
    contract_address: String,
    mint_amount: String,
    receipt_address: String,
) -> Result<Response<QcoreMsg>, PoolError> {

    let msg_mint_tribute = QcoreMsg::msg_mint_tribute(creator, contract_address, mint_amount, receipt_address);

    let res = Response::new()
        .add_attribute("method", "msg_mint_tribute")
        .add_message(msg_mint_tribute);

    Ok(res)
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(deps: Deps<QcoreQuery>, _env: Env, msg: QueryMsg) -> StdResult<Binary> {

    match msg {
        QueryMsg::QueryBlockEmissionRequest { block_number, } => to_json_binary(&query_block_emission_request(deps, block_number)?),
    }
}

fn query_block_emission_request(deps: Deps<QcoreQuery>, block_number: String) -> StdResult<QueryBlockEmissionResponse> {
    info!("contract-pool-query_block_emission_request: block_number {}  Mb/s", block_number);
    let querier = QcoreQuerier::new(&deps.querier);

    let response = querier.query_block_emission_request(block_number).unwrap();

    info!("contract-pool-query_block_emission_request: response {}  Mb/s", response.block_emission);

    Ok(QueryBlockEmissionResponse {
        block_emission: response.block_emission,
    })
}