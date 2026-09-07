import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CycleCountService } from '../../../services/CycleCount.service';
import { CycleCount } from '../../../models/CycleCount';
import { SubBaseComponent } from '../../CycleCount/sub.base.component';

@Component({
    selector: 'app-create-cycleCount',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCycleCountComponent extends SubBaseComponent implements OnInit {

    title = 'Add CycleCount';

    cycleCountForm: FormGroup;
    cycleCount: CycleCount;

    constructor( http: HttpClient,
        private cycleCountService: CycleCountService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status): void {
        this.cycleCountService
        .addCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCycleCount']);
            });
    }

    ngOnInit(): void {
    }
}