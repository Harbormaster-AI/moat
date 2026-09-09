import React, { Component } from 'react'
import QualitySpecificationService from '../services/QualitySpecificationService'

class ViewQualitySpecificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            qualitySpecification: {}
        }
    }

    componentDidMount(){
        QualitySpecificationService.getQualitySpecificationById(this.state.id).then( res => {
            this.setState({qualitySpecification: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View QualitySpecification Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> specCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.qualitySpecification.specCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.qualitySpecification.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> version:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.qualitySpecification.version }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewQualitySpecificationComponent
