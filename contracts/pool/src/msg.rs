use cosmwasm_schema::{cw_serde, QueryResponses};

#[cw_serde]
pub struct InstantiateMsg {
    pub block_number: String,
}


#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {

    #[returns(QueryBlockEmissionResponse)]
    QueryBlockEmissionRequest { block_number: String },

}

// We define a custom struct for each query response
#[cw_serde]
pub struct QueryBlockEmissionResponse {
    pub block_emission: String,
}

#[cw_serde]
pub enum ExecuteMsg {
    MsgMintTribute {
        creator:         String,
        contract_address: String,
        mint_amount:      String,
        receipt_address:  String,
    },
}