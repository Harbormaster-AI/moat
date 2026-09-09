import React, { Component } from 'react'
import UoMConversionService from '../services/UoMConversionService'

class ViewUoMConversionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            uoMConversion: {}
        }
    }

    componentDidMount(){
        UoMConversionService.getUoMConversionById(this.state.id).then( res => {
            this.setState({uoMConversion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View UoMConversion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> factor:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.uoMConversion.factor }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> precision:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.uoMConversion.precision }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FromUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.uoMConversion.fromUnit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ToUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.uoMConversion.toUnit }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewUoMConversionComponent
