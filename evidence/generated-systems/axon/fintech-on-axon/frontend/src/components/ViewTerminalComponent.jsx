import React, { Component } from 'react'
import TerminalService from '../services/TerminalService'

class ViewTerminalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            terminal: {}
        }
    }

    componentDidMount(){
        TerminalService.getTerminalById(this.state.id).then( res => {
            this.setState({terminal: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Terminal Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> location:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.terminal.location }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.terminal.type }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.terminal.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTerminalComponent
