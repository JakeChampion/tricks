import { getStore } from "@netlify/blobs";

let r = [];
let i = [];
async function f(request, init) {
    r.push(request) 
    i.push(init)
    console.log({request, init})
    return fetch(request, init)
}

export default async (req, context) => {
    const construction = getStore({ fetch: f, name: "construction" });
    await construction.set("nails", 9);
    const entry = await construction.get("nails")
    return Response.json({r,getStore:getStore.toString(),set: construction.set.toString(),get: construction.get.toString()});
};

export const config = {
    path: '/f'
};
