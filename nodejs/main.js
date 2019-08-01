const circle = require('./libs/circle');

exports.handler = function(context, event) {
    context.callback(circle.area(6));
};
