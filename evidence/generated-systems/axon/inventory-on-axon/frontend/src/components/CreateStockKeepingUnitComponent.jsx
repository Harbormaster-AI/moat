import React, { Component } from 'react'
import StockKeepingUnitService from '../services/StockKeepingUnitService';

class CreateStockKeepingUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                skuCode: '',
                name: '',
                weight: '',
                weightUnit: '',
                volume: '',
                volumeUnit: '',
                shelfLifeDays: '',
                hazardousMaterial: '',
                itemType: '',
                unitOfMeasure: ''
        }
        this.changeskuCodeHandler = this.changeskuCodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeweightHandler = this.changeweightHandler.bind(this);
        this.changeweightUnitHandler = this.changeweightUnitHandler.bind(this);
        this.changevolumeHandler = this.changevolumeHandler.bind(this);
        this.changevolumeUnitHandler = this.changevolumeUnitHandler.bind(this);
        this.changeshelfLifeDaysHandler = this.changeshelfLifeDaysHandler.bind(this);
        this.changehazardousMaterialHandler = this.changehazardousMaterialHandler.bind(this);
        this.changeItemTypeHandler = this.changeItemTypeHandler.bind(this);
        this.changeUnitOfMeasureHandler = this.changeUnitOfMeasureHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            StockKeepingUnitService.getStockKeepingUnitById(this.state.id).then( (res) =>{
                let stockKeepingUnit = res.data;
                this.setState({
                    skuCode: stockKeepingUnit.skuCode,
                    name: stockKeepingUnit.name,
                    weight: stockKeepingUnit.weight,
                    weightUnit: stockKeepingUnit.weightUnit,
                    volume: stockKeepingUnit.volume,
                    volumeUnit: stockKeepingUnit.volumeUnit,
                    shelfLifeDays: stockKeepingUnit.shelfLifeDays,
                    hazardousMaterial: stockKeepingUnit.hazardousMaterial,
                    itemType: stockKeepingUnit.itemType,
                    unitOfMeasure: stockKeepingUnit.unitOfMeasure
                });
            });
        }        
    }
    saveOrUpdateStockKeepingUnit = (e) => {
        e.preventDefault();
        let stockKeepingUnit = {
                stockKeepingUnitId: this.state.id,
                skuCode: this.state.skuCode,
                name: this.state.name,
                weight: this.state.weight,
                weightUnit: this.state.weightUnit,
                volume: this.state.volume,
                volumeUnit: this.state.volumeUnit,
                shelfLifeDays: this.state.shelfLifeDays,
                hazardousMaterial: this.state.hazardousMaterial,
                itemType: this.state.itemType,
                unitOfMeasure: this.state.unitOfMeasure
            };
        console.log('stockKeepingUnit => ' + JSON.stringify(stockKeepingUnit));

        // step 5
        if(this.state.id === '_add'){
            stockKeepingUnit.stockKeepingUnitId=''
            StockKeepingUnitService.createStockKeepingUnit(stockKeepingUnit).then(res =>{
                this.props.history.push('/stockKeepingUnits');
            });
        }else{
            StockKeepingUnitService.updateStockKeepingUnit(stockKeepingUnit).then( res => {
                this.props.history.push('/stockKeepingUnits');
            });
        }
    }
    
    changeskuCodeHandler= (event) => {
        this.setState({skuCode: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeweightHandler= (event) => {
        this.setState({weight: event.target.value});
    }
    changeweightUnitHandler= (event) => {
        this.setState({weightUnit: event.target.value});
    }
    changevolumeHandler= (event) => {
        this.setState({volume: event.target.value});
    }
    changevolumeUnitHandler= (event) => {
        this.setState({volumeUnit: event.target.value});
    }
    changeshelfLifeDaysHandler= (event) => {
        this.setState({shelfLifeDays: event.target.value});
    }
    changehazardousMaterialHandler= (event) => {
        this.setState({hazardousMaterial: event.target.value});
    }
    changeItemTypeHandler= (event) => {
        this.setState({itemType: event.target.value});
    }
    changeUnitOfMeasureHandler= (event) => {
        this.setState({unitOfMeasure: event.target.value});
    }

    cancel(){
        this.props.history.push('/stockKeepingUnits');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add StockKeepingUnit</h3>
        }else{
            return <h3 className="text-center">Update StockKeepingUnit</h3>
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
                                            <label> skuCode:&emsp; </label>
                                                <input placeholder="skuCode" name="skuCode" className="form-control" value={this.state.skuCode} onChange={this.changeskuCodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> weight:&emsp; </label>
                                                <input placeholder="weight" name="weight" className="form-control" value={this.state.weight} onChange={this.changeweightHandler}/>

                                            <label> weightUnit:&emsp; </label>
                                                <input placeholder="weightUnit" name="weightUnit" className="form-control" value={this.state.weightUnit} onChange={this.changeweightUnitHandler}/>

                                            <label> volume:&emsp; </label>
                                                <input placeholder="volume" name="volume" className="form-control" value={this.state.volume} onChange={this.changevolumeHandler}/>

                                            <label> volumeUnit:&emsp; </label>
                                                <input placeholder="volumeUnit" name="volumeUnit" className="form-control" value={this.state.volumeUnit} onChange={this.changevolumeUnitHandler}/>

                                            <label> shelfLifeDays:&emsp; </label>
                                                <input type="number" placeholder="shelfLifeDays" name="shelfLifeDays" className="form-control" value={this.state.shelfLifeDays} onChange={this.changeshelfLifeDaysHandler}/>

                                            <label> hazardousMaterial:&emsp; </label>
                                                <input type="checkbox" placeholder="hazardousMaterial" name="hazardousMaterial" className="form-control" value={this.state.hazardousMaterial} onChange={this.changehazardousMaterialHandler}/>


                                            <label> ItemType:&emsp; </label>
                                                <select value={this.state.itemType} onChange={this.changeItemTypeHandler}>
                      <option name="ItemType" className="form-control" >
                          FinishedGood
                      </option>
                      <option name="ItemType" className="form-control" >
                          Component
                      </option>
                      <option name="ItemType" className="form-control" >
                          RawMaterial
                      </option>
                      <option name="ItemType" className="form-control" >
                          Packaging
                      </option>
                      <option name="ItemType" className="form-control" >
                          SparePart
                      </option>
                      <option name="ItemType" className="form-control" >
                          Consumable
                      </option>
                    </select>

                                            <label> UnitOfMeasure:&emsp; </label>
                                                <select value={this.state.unitOfMeasure} onChange={this.changeUnitOfMeasureHandler}>
                      <option name="UnitOfMeasure" className="form-control" >
                          Each
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Case
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pallet
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Dozen
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Gram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Kilogram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pound
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Ounce
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Milliliter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Liter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          CubicMeter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Meter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Foot
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          SquareMeter
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateStockKeepingUnit}>Save</button>
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

export default CreateStockKeepingUnitComponent
