import React, { Component } from 'react'
import DataCategoryService from '../services/DataCategoryService'

class ViewDataCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataCategory: {}
        }
    }

    componentDidMount(){
        DataCategoryService.getDataCategoryById(this.state.id).then( res => {
            this.setState({dataCategory: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataCategory Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataCategory.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataCategory.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Classification:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataCategory.classification }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataCategoryComponent
