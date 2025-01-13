const { getStore } = require("@netlify/blobs");

let r = [];
let i = [];
async function f(request, init) {
    r.push(request) 
    i.push(init)
    console.log({request, init})
    return fetch(request, init)
}

module.exports.handler = async (req, context) => {
    const construction = getStore("construction");
    await construction.set("nails", 9);
    const entry = await construction.get("nails")
    return Response.json({r,getStore:getStore.toString(),set: construction.set.toString(),get: construction.get.toString()});
};
