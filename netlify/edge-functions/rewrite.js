export default async (request, context) => {
    return new Response("howdy");
};

export const config = {
    path: '/*'
};
