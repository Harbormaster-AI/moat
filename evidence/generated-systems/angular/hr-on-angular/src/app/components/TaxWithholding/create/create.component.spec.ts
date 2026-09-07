
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTaxWithholdingComponent } from './create.component';
import { TaxWithholdingService } from '../../../services/TaxWithholding.service';
import { Router } from '@angular/router';

describe('CreateTaxWithholdingComponent', () => {
  let component: CreateTaxWithholdingComponent;
  let fixture: ComponentFixture<CreateTaxWithholdingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTaxWithholdingComponent
      ],
      providers: [
        TaxWithholdingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTaxWithholdingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});