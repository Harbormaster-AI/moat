import React, { Component } from 'react'
import ProductionLineService from '../services/ProductionLineService';

class UpdateProductionLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                lineType: ''
        }
        this.updateProductionLine = this.updateProductionLine.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeLineTypeHandler = this.changeLineTypeHandler.bind(this);
    }

    componentDidMount(){
        ProductionLineService.getProductionLineById(this.state.id).then( (res) =>{
            let productionLine = res.data;
            this.setState({
                name: productionLine.name,
                lineType: productionLine.lineType
            });
        });
    }

    updateProductionLine = (e) => {
        e.preventDefault();
        let productionLine = {
            productionLineId: this.state.id,
            name: this.state.name,
            lineType: this.state.lineType
        };
        console.log('productionLine => ' + JSON.stringify(productionLine));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductionLineService.updateProductionLine(productionLine).then( res => {
            this.props.history.push('/productionLines');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeLineTypeHandler= (event) => {
        this.setState({lineType: event.target.value});
    }

    cancel(){
        this.props.history.push('/productionLines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductionLine</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> LineType: </label>
                                                <select value={this.state.lineType} onChange={this.changeLineTypeHandler}>
                      <option name="LineType" className="form-control" >
                          FinalAssembly
                      </option>
                      <option name="LineType" className="form-control" >
                          SubAssembly
                      </option>
                      <option name="LineType" className="form-control" >
                          Integration
                      </option>
                      <option name="LineType" className="form-control" >
                          TestAndDelivery
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductionLine}>Save</button>
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

export default UpdateProductionLineComponent
