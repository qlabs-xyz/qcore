use cosmwasm_std::{QuerierWrapper, QueryRequest, StdResult};

use crate::query::{
    QueryBlockEmissionResponse, QcoreQuery,
};

/// This is a helper wrapper to easily use our custom queries
pub struct QcoreQuerier<'a> {
    querier: &'a QuerierWrapper<'a, QcoreQuery>,
}

impl<'a> QcoreQuerier<'a> {
    pub fn new(querier: &'a QuerierWrapper<QcoreQuery>) -> Self {
        QcoreQuerier { querier }
    }

    pub fn query_block_emission_request(
        &self,
        block_number: String,
    ) -> StdResult<QueryBlockEmissionResponse> {
        let query_block_emission_request = QcoreQuery::QueryBlockEmissionRequest {
            block_number,
        };
        let request: QueryRequest<QcoreQuery> = QcoreQuery::into(query_block_emission_request);
        self.querier.query(&request)
    }
}



