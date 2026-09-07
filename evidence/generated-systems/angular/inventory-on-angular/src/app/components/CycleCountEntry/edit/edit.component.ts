import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';
import { SubBaseComponent } from '../../CycleCountEntry/sub.base.component';


@Component({
    selector: 'app-edit-cycleCountEntry',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCycleCountEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CycleCountEntry';

    cycleCountEntryForm: FormGroup;
    cycleCountEntry: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CycleCountEntryService,
        private fb: FormBuilder
) {
        super(http);
        this.cycleCountEntryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      systemQuantity: ['', Validators.required],
      countedQuantity: ['', Validators.required],
      varianceQuantity: ['', Validators.required],
      recountRequired: ['', Validators.required],
      CycleCount: ['', ],
      Sku: ['', ],
      Lot: ['', ],
      Location: ['', ],
      SerialNumbers: ['', ],
      StockStatus: ['', ]
        });
    }

    
    updateCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCycleCountEntry']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCycleCountEntry(params['id']).subscribe(res => {
                this.cycleCountEntry = res;
            });
        });
    }
}