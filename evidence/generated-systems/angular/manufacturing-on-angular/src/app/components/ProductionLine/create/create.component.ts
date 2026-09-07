import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductionLineService } from '../../../services/ProductionLine.service';
import { ProductionLine } from '../../../models/ProductionLine';
import { SubBaseComponent } from '../../ProductionLine/sub.base.component';

@Component({
    selector: 'app-create-productionLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductionLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProductionLine';

    productionLineForm: FormGroup;
    productionLine: ProductionLine;

    constructor( http: HttpClient,
        private productionLineService: ProductionLineService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.productionLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      lineCode: ['', Validators.required],
      Plant: ['', ],
      WorkCenters: ['', ],
      LineType: ['', ]
        });
    }

    
    addProductionLine(name, lineCode, Plant, WorkCenters, LineType): void {
        this.productionLineService
        .addProductionLine(name, lineCode, Plant, WorkCenters, LineType)
            .subscribe(() => {
                this.router.navigate(['/indexProductionLine']);
            });
    }

    ngOnInit(): void {
    }
}