import React, { Component } from 'react'
import PriceBookService from '../services/PriceBookService'

class ViewPriceBookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            priceBook: {}
        }
    }

    componentDidMount(){
        PriceBookService.getPriceBookById(this.state.id).then( res => {
            this.setState({priceBook: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PriceBook Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBook.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBook.asActive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.priceBook.description }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPriceBookComponent
