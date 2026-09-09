import React, { Component } from 'react'
import DataProviderService from '../services/DataProviderService'

class ViewDataProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataProvider: {}
        }
    }

    componentDidMount(){
        DataProviderService.getDataProviderById(this.state.id).then( res => {
            this.setState({dataProvider: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataProvider Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProvider.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProvider.website }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProviderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProvider.providerType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataProviderComponent
