import gql from 'graphql-tag';

export const QUERY_JOB_LIST = gql`
    query Jobs {
        jobs(input: {$type: String}) {
            id,
            title,
            slug
        }
    }
`;
