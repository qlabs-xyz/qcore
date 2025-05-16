use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::CustomQuery;

#[cw_serde]
#[derive(QueryResponses)]
pub enum QcoreQuery {

    #[returns(QueryBlockEmissionResponse)]
    QueryBlockEmissionRequest {
        block_number: String,
    },
}

impl CustomQuery for QcoreQuery {}

impl QcoreQuery {
}

#[cw_serde]
pub struct QueryBlockEmissionResponse {
    pub block_emission: String,
}