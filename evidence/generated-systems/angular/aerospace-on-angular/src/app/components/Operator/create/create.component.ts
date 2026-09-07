import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OperatorService } from '../../../services/Operator.service';
import { Operator } from '../../../models/Operator';
import { SubBaseComponent } from '../../Operator/sub.base.component';

@Component({
    selector: 'app-create-operator',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOperatorComponent extends SubBaseComponent implements OnInit {

    title = 'Add Operator';

    operatorForm: FormGroup;
    operator: Operator;

    constructor( http: HttpClient,
        private operatorService: OperatorService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.operatorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      icaoDesignator: ['', Validators.required],
      AircraftOrders: ['', ],
      OperatedAircraft: ['', ],
      SalesRegion: ['', ],
      OperatorType: ['', ]
        });
    }

    
    addOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType): void {
        this.operatorService
        .addOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType)
            .subscribe(() => {
                this.router.navigate(['/indexOperator']);
            });
    }

    ngOnInit(): void {
    }
}