import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { NonconformanceService } from '../../../services/Nonconformance.service';
import { SubBaseComponent } from '../../Nonconformance/sub.base.component';


@Component({
    selector: 'app-edit-nonconformance',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditNonconformanceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Nonconformance';

    nonconformanceForm: FormGroup;
    nonconformance: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: NonconformanceService,
        private fb: FormBuilder
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

    
    updateNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexNonconformance']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getNonconformance(params['id']).subscribe(res => {
                this.nonconformance = res;
            });
        });
    }
}