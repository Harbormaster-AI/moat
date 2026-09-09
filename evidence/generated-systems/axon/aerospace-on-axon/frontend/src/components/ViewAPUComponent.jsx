import React, { Component } from 'react'
import APUService from '../services/APUService'

class ViewAPUComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aPU: {}
        }
    }

    componentDidMount(){
        APUService.getAPUById(this.state.id).then( res => {
            this.setState({aPU: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View APU Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> model:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aPU.model }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAPUComponent
