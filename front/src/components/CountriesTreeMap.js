import React from 'react';
import { rd3, Treemap } from 'react-d3';
import TransactionStore from '../stores/TransactionStore';

export default class CountriesTreeMap extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            'countries': this.props.countries
        };
        this.change = this.change.bind(this);
    }

    render() {
        return (
            <Treemap data={this.state.countries}
             width={500} height={500}
              textColor='#3182bd' fontSize="12px" title='Treemap countries transactions count'/>
        );
    }


    change() {
        this.setState({
            "countries": TransactionStore.getCountryCount()
        });
    }

    componentDidMount() {
        TransactionStore.addChangeListener(this.change);
    }

    componentWillUnmount() {
        TransactionStore.removeChangeListener(this.change);
    }
}

 CountriesTreeMap.defaultProps = { 'countries':[] };