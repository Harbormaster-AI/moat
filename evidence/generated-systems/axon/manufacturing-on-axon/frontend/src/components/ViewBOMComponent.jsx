import React, { Component } from 'react'
import BOMService from '../services/BOMService'

class ViewBOMComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            bOM: {}
        }
    }

    componentDidMount(){
        BOMService.getBOMById(this.state.id).then( res => {
            this.setState({bOM: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BOM Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bomNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOM.bomNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> revision:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOM.revision }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivityStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOM.effectivityStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivityEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOM.effectivityEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.bOM.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBOMComponent
