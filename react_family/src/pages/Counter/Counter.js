import React, {Component} from 'react';
import {increment, decrement, reset} from 'actions/counter';
import {connect} from 'react-redux';
/*
 * Counter组件和Redux联合起来，使Counter能获得Redux的state，并且能发射action
 * connect接收两个参数，一个mapStateToProps，就是把redux的state转为组件的Props，还有一个参数mapDispatchToProps，就是把
 * 发射actions的方法转为Props属性函数。connect函数作用是从Redux state树中读取部分数据，并通过props来吧这些数据提供给要渲染
 * 的组件，也传递dispach(action)函数到props
 */
class Counter extends Component {
    render() {
        return(
            <div>
                <div>当前计数为{this.props.counter.count}</div>
                <button onClick={() => this.props.increment()}>自增
                </button>
                <button onClick={() => this.props.decrement()}>自减
                </button>
                <button onClick={() => this.props.reset()}>重置
                </button>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        counter: state.counter
    }
};

const mapDispatchToProps = (dispatch) => {
    return {
        increment: () => {
            dispatch(increment())
        },
        decrement: () => {
            dispatch(decrement())
        },
        reset: () => {
            dispatch(reset())
        }
    }
};
export default connect(mapStateToProps, mapDispatchToProps)(Counter);





















