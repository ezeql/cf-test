import React from 'react/addons';
import {FixedDataTable,Table,Column} from 'fixed-data-table';

import TransactionStore from '../stores/TransactionStore';
import TransactionService from '../services/TransactionService';

require('../styles/fixed-data-table.min.css');
require('../styles/bootstrap.min.css');

//TransactionService.start();

export default class List extends React.Component{

   constructor(props) {
        super(props);
        this.state = {
             'rows': this.props.rows
        };
        this.change = this.change.bind(this);
    }

    rowGetter(row){
        const data = this.state.rows[row];
        return data !== undefined ? data.data : {};

    }

    render(){

            return (
        <Table  rowHeight={50} rowGetter={ row => this.state.rows[row].data } rowsCount={this.state.rows.length} 
            width={1000} height={500} headerHeight={50}>
            <Column   label="userId"              width={125}   dataKey={"userId"}/>
            <Column   label="currencyFrom"        width={50}   dataKey={"currencyFrom"}/>
            <Column   label="currencyTo"          width={50}   dataKey={"currencyTo"}/>
            <Column   label="amountSell"          width={125}   dataKey={"amountSell"}/>
            <Column   label="amountBuy"           width={125}   dataKey={"amountBuy"}/>
            <Column   label="rate"                width={125}   dataKey={"rate"}/>
            <Column   label="timePlaced"          width={200}   dataKey={"timePlaced"}/>
            <Column   label="originatingCountry"  width={200}   dataKey={"originatingCountry"}/>
        </Table>);

    }
    change(){
        this.setState({
            "rows": TransactionStore.getAll()
        });

    }


    componentDidMount() {
        TransactionStore.addChangeListener(this.change);
    }

    componentWillUnmount(){
        TransactionStore.removeChangeListener(this.change);
    }




}

 List.defaultProps = { 'rows':[] };