import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NonconformanceService } from '../../../services/Nonconformance.service';
import { Nonconformance } from '../../../models/Nonconformance';
import { SubBaseComponent } from '../../Nonconformance/sub.base.component';

@Component({
    selector: 'app-create-nonconformance',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateNonconformanceComponent extends SubBaseComponent implements OnInit {

    title = 'Add Nonconformance';

    nonconformanceForm: FormGroup;
    nonconformance: Nonconformance;

    constructor( http: HttpClient,
        private nonconformanceService: NonconformanceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.nonconformanceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  ncNumber: ['', Validators.required],
      description: ['', Validators.required],
      containmentAction: ['', Validators.required],
      Item: ['', ],
      WorkOrder: ['', ],
      InspectionLot: ['', ],
      CorrectiveAction: ['', ],
      NcType: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    addNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status): void {
        this.nonconformanceService
        .addNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status)
            .subscribe(() => {
                this.router.navigate(['/indexNonconformance']);
            });
    }

    ngOnInit(): void {
    }
}