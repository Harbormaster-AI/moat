import React, { Component } from 'react'
import CategoryService from '../services/CategoryService'

class ViewCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            category: {}
        }
    }

    componentDidMount(){
        CategoryService.getCategoryById(this.state.id).then( res => {
            this.setState({category: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Category Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.category.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> slug:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.category.slug }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> position:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.category.position }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.category.asActive }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCategoryComponent
