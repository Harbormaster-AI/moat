import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';
import { CycleCountEntry } from '../../../models/CycleCountEntry';
import { SubBaseComponent } from '../../CycleCountEntry/sub.base.component';

@Component({
    selector: 'app-create-cycleCountEntry',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCycleCountEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Add CycleCountEntry';

    cycleCountEntryForm: FormGroup;
    cycleCountEntry: CycleCountEntry;

    constructor( http: HttpClient,
        private cycleCountEntryService: CycleCountEntryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus): void {
        this.cycleCountEntryService
        .addCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus)
            .subscribe(() => {
                this.router.navigate(['/indexCycleCountEntry']);
            });
    }

    ngOnInit(): void {
    }
}