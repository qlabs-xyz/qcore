use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use cosmwasm_schema::cw_serde;
use cosmwasm_std::{CosmosMsg, CustomMsg};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    QueryBlockEmissionRequest { block_number: String },
}

#[cw_serde]
pub enum QcoreMsg {
    MsgMintTribute {
        creator: String,
        contract_address: String,
        mint_amount: String,
        receipt_address: String,
    },
}

impl QcoreMsg {
    pub fn msg_mint_tribute(creator: String, contract_address: String, mint_amount: String, receipt_address: String) -> Self {
        QcoreMsg::MsgMintTribute {
            creator,
            contract_address,
            mint_amount,
            receipt_address,
        }
    }
}

impl From<QcoreMsg> for CosmosMsg<QcoreMsg> {
    fn from(msg: QcoreMsg) -> CosmosMsg<QcoreMsg> {
        CosmosMsg::Custom(msg)
    }
}

impl CustomMsg for QcoreMsg {}