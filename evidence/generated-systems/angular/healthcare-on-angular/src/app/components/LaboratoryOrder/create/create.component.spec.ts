
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLaboratoryOrderComponent } from './create.component';
import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';
import { Router } from '@angular/router';

describe('CreateLaboratoryOrderComponent', () => {
  let component: CreateLaboratoryOrderComponent;
  let fixture: ComponentFixture<CreateLaboratoryOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLaboratoryOrderComponent
      ],
      providers: [
        LaboratoryOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLaboratoryOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});