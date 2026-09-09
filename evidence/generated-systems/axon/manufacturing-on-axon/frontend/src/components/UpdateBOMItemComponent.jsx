import React, { Component } from 'react'
import BOMItemService from '../services/BOMItemService';

class UpdateBOMItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                scrapPercent: ''
        }
        this.updateBOMItem = this.updateBOMItem.bind(this);

        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changescrapPercentHandler = this.changescrapPercentHandler.bind(this);
    }

    componentDidMount(){
        BOMItemService.getBOMItemById(this.state.id).then( (res) =>{
            let bOMItem = res.data;
            this.setState({
                lineNumber: bOMItem.lineNumber,
                quantity: bOMItem.quantity,
                scrapPercent: bOMItem.scrapPercent
            });
        });
    }

    updateBOMItem = (e) => {
        e.preventDefault();
        let bOMItem = {
            bOMItemId: this.state.id,
            lineNumber: this.state.lineNumber,
            quantity: this.state.quantity,
            scrapPercent: this.state.scrapPercent
        };
        console.log('bOMItem => ' + JSON.stringify(bOMItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        BOMItemService.updateBOMItem(bOMItem).then( res => {
            this.props.history.push('/bOMItems');
        });
    }

    changelineNumberHandler= (event) => {
        this.setState({lineNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changescrapPercentHandler= (event) => {
        this.setState({scrapPercent: event.target.value});
    }

    cancel(){
        this.props.history.push('/bOMItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BOMItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber: </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> scrapPercent: </label>
                                                <input placeholder="scrapPercent" name="scrapPercent" className="form-control" value={this.state.scrapPercent} onChange={this.changescrapPercentHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBOMItem}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateBOMItemComponent
