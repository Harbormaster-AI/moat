import React, { Component } from 'react'
import ProductionLineService from '../services/ProductionLineService';

class CreateProductionLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                lineCode: '',
                lineType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelineCodeHandler = this.changelineCodeHandler.bind(this);
        this.changeLineTypeHandler = this.changeLineTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductionLineService.getProductionLineById(this.state.id).then( (res) =>{
                let productionLine = res.data;
                this.setState({
                    name: productionLine.name,
                    lineCode: productionLine.lineCode,
                    lineType: productionLine.lineType
                });
            });
        }        
    }
    saveOrUpdateProductionLine = (e) => {
        e.preventDefault();
        let productionLine = {
                productionLineId: this.state.id,
                name: this.state.name,
                lineCode: this.state.lineCode,
                lineType: this.state.lineType
            };
        console.log('productionLine => ' + JSON.stringify(productionLine));

        // step 5
        if(this.state.id === '_add'){
            productionLine.productionLineId=''
            ProductionLineService.createProductionLine(productionLine).then(res =>{
                this.props.history.push('/productionLines');
            });
        }else{
            ProductionLineService.updateProductionLine(productionLine).then( res => {
                this.props.history.push('/productionLines');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelineCodeHandler= (event) => {
        this.setState({lineCode: event.target.value});
    }
    changeLineTypeHandler= (event) => {
        this.setState({lineType: event.target.value});
    }

    cancel(){
        this.props.history.push('/productionLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductionLine</h3>
        }else{
            return <h3 className="text-center">Update ProductionLine</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> lineCode:&emsp; </label>
                                                <input placeholder="lineCode" name="lineCode" className="form-control" value={this.state.lineCode} onChange={this.changelineCodeHandler}/>

                                            <label> LineType:&emsp; </label>
                                                <select value={this.state.lineType} onChange={this.changeLineTypeHandler}>
                      <option name="LineType" className="form-control" >
                          Discrete
                      </option>
                      <option name="LineType" className="form-control" >
                          Batch
                      </option>
                      <option name="LineType" className="form-control" >
                          Continuous
                      </option>
                      <option name="LineType" className="form-control" >
                          FlexibleCell
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductionLine}>Save</button>
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

export default CreateProductionLineComponent
