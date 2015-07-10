import React from 'react';
import TransactionStore from '../stores/TransactionStore';

export default class Graphs extends React.Component { 
    constructor(props) {
        super(props);
        this.state = { };
        this.change = this.change.bind(this);
    }

    render() {
        return (
            <div>TBD</div>
        );
    }


    change() {
        // this.setState({
        //     
        // });
    }

    componentDidMount() {
        TransactionStore.addChangeListener(this.change);
    }

    componentWillUnmount() {
        TransactionStore.removeChangeListener(this.change);
    }
}

 Graphs.defaultProps = { };