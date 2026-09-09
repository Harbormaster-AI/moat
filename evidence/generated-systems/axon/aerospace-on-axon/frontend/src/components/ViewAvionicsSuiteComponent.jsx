import React, { Component } from 'react'
import AvionicsSuiteService from '../services/AvionicsSuiteService'

class ViewAvionicsSuiteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            avionicsSuite: {}
        }
    }

    componentDidMount(){
        AvionicsSuiteService.getAvionicsSuiteById(this.state.id).then( res => {
            this.setState({avionicsSuite: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AvionicsSuite Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> suiteName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.avionicsSuite.suiteName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> softwareBaseline:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.avionicsSuite.softwareBaseline }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAvionicsSuiteComponent
