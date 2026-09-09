import React, { Component } from 'react'
import FeatureService from '../services/FeatureService'

class ViewFeatureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            feature: {}
        }
    }

    componentDidMount(){
        FeatureService.getFeatureById(this.state.id).then( res => {
            this.setState({feature: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Feature Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feature.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feature.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DataType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feature.dataType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFeatureComponent
