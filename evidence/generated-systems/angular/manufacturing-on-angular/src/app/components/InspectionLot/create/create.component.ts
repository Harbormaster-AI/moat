import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InspectionLotService } from '../../../services/InspectionLot.service';
import { InspectionLot } from '../../../models/InspectionLot';
import { SubBaseComponent } from '../../InspectionLot/sub.base.component';

@Component({
    selector: 'app-create-inspectionLot',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInspectionLotComponent extends SubBaseComponent implements OnInit {

    title = 'Add InspectionLot';

    inspectionLotForm: FormGroup;
    inspectionLot: InspectionLot;

    constructor( http: HttpClient,
        private inspectionLotService: InspectionLotService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status): void {
        this.inspectionLotService
        .addInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInspectionLot']);
            });
    }

    ngOnInit(): void {
    }
}