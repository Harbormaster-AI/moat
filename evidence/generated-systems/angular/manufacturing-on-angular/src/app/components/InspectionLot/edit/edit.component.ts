import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InspectionLotService } from '../../../services/InspectionLot.service';
import { SubBaseComponent } from '../../InspectionLot/sub.base.component';


@Component({
    selector: 'app-edit-inspectionLot',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInspectionLotComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InspectionLot';

    inspectionLotForm: FormGroup;
    inspectionLot: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InspectionLotService,
        private fb: FormBuilder
) {
        super(http);
        this.inspectionLotForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lotNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      sampleSize: ['', Validators.required],
      createdOn: ['', Validators.required],
      Item: ['', ],
      WorkOrder: ['', ],
      GoodsReceipt: ['', ],
      Results: ['', ],
      InspectionType: ['', ],
      Status: ['', ]
        });
    }

    
    updateInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInspectionLot']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInspectionLot(params['id']).subscribe(res => {
                this.inspectionLot = res;
            });
        });
    }
}