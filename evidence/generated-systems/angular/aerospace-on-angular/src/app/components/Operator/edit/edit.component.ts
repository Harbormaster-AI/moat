import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OperatorService } from '../../../services/Operator.service';
import { SubBaseComponent } from '../../Operator/sub.base.component';


@Component({
    selector: 'app-edit-operator',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOperatorComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Operator';

    operatorForm: FormGroup;
    operator: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OperatorService,
        private fb: FormBuilder
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

    
    updateOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOperator']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOperator(params['id']).subscribe(res => {
                this.operator = res;
            });
        });
    }
}