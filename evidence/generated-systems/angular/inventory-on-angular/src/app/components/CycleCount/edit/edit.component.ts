import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CycleCountService } from '../../../services/CycleCount.service';
import { SubBaseComponent } from '../../CycleCount/sub.base.component';


@Component({
    selector: 'app-edit-cycleCount',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCycleCountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CycleCount';

    cycleCountForm: FormGroup;
    cycleCount: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CycleCountService,
        private fb: FormBuilder
) {
        super(http);
        this.cycleCountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  countNumber: ['', Validators.required],
      scheduledDate: ['', Validators.required],
      performedDate: ['', Validators.required],
      approvedBy: ['', Validators.required],
      Warehouse: ['', ],
      Locations: ['', ],
      Entries: ['', ],
      Transactions: ['', ],
      Status: ['', ]
        });
    }

    
    updateCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCycleCount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCycleCount(params['id']).subscribe(res => {
                this.cycleCount = res;
            });
        });
    }
}