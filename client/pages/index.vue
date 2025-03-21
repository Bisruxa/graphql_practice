<template>
  <div>
  <p>There are {{ data?.Books?.length || 0 }} Books.</p>
  <ul>
    <li v-for="Book in data?.Books" :key="Book.id">
      📖 {{ Book.title }} 
    </li>
  </ul>
  </div>
</template>

<script setup>
const query = gql`
  query getBooks($limit: Int!) {
    Books(limit: $limit) {
      uuid
      title
      
    }
  }
`;

const { result: data } = useQuery(query, { limit: 5 });
</script>
