import counter from 'reducers/counter';
import userInfo from 'reducers/userInfo';
//组合reducer
export default function combineReducers(state = {}, action) {
    return {
        counter: counter(state.counter, action),
        userInfo: userInfo(state.userInfo, action)
    }
}