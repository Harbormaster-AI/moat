import React, { Component } from 'react'
import Component_Service from '../services/Component_Service'

class ViewComponent_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            component_: {}
        }
    }

    componentDidMount(){
        Component_Service.getComponent_ById(this.state.id).then( res => {
            this.setState({component_: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Component_ Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> partNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.component_.partNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.component_.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ComponentCategory:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.component_.componentCategory }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SerializationMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.component_.serializationMethod }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewComponent_Component
