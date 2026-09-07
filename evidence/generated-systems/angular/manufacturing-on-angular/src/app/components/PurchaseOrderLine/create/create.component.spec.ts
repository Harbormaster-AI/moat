
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePurchaseOrderLineComponent } from './create.component';
import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';
import { Router } from '@angular/router';

describe('CreatePurchaseOrderLineComponent', () => {
  let component: CreatePurchaseOrderLineComponent;
  let fixture: ComponentFixture<CreatePurchaseOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePurchaseOrderLineComponent
      ],
      providers: [
        PurchaseOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePurchaseOrderLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});