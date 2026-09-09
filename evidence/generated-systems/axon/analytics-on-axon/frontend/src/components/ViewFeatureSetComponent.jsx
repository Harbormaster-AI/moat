import React, { Component } from 'react'
import FeatureSetService from '../services/FeatureSetService'

class ViewFeatureSetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            featureSet: {}
        }
    }

    componentDidMount(){
        FeatureSetService.getFeatureSetById(this.state.id).then( res => {
            this.setState({featureSet: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FeatureSet Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.featureSet.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> refreshSchedule:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.featureSet.refreshSchedule }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> StoreType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.featureSet.storeType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFeatureSetComponent
