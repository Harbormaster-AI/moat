import React, { Component } from 'react'
import ContentCategoryService from '../services/ContentCategoryService'

class ViewContentCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            contentCategory: {}
        }
    }

    componentDidMount(){
        ContentCategoryService.getContentCategoryById(this.state.id).then( res => {
            this.setState({contentCategory: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ContentCategory Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contentCategory.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contentCategory.name }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewContentCategoryComponent
