import React, { Component } from 'react'
import BOMItemService from '../services/BOMItemService';

class CreateBOMItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                scrapPercent: ''
        }
        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changescrapPercentHandler = this.changescrapPercentHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BOMItemService.getBOMItemById(this.state.id).then( (res) =>{
                let bOMItem = res.data;
                this.setState({
                    lineNumber: bOMItem.lineNumber,
                    quantity: bOMItem.quantity,
                    scrapPercent: bOMItem.scrapPercent
                });
            });
        }        
    }
    saveOrUpdateBOMItem = (e) => {
        e.preventDefault();
        let bOMItem = {
                bOMItemId: this.state.id,
                lineNumber: this.state.lineNumber,
                quantity: this.state.quantity,
                scrapPercent: this.state.scrapPercent
            };
        console.log('bOMItem => ' + JSON.stringify(bOMItem));

        // step 5
        if(this.state.id === '_add'){
            bOMItem.bOMItemId=''
            BOMItemService.createBOMItem(bOMItem).then(res =>{
                this.props.history.push('/bOMItems');
            });
        }else{
            BOMItemService.updateBOMItem(bOMItem).then( res => {
                this.props.history.push('/bOMItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BOMItem</h3>
        }else{
            return <h3 className="text-center">Update BOMItem</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber:&emsp; </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> scrapPercent:&emsp; </label>
                                                <input placeholder="scrapPercent" name="scrapPercent" className="form-control" value={this.state.scrapPercent} onChange={this.changescrapPercentHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBOMItem}>Save</button>
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

export default CreateBOMItemComponent
