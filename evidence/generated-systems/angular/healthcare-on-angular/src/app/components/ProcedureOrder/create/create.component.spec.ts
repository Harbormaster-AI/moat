
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProcedureOrderComponent } from './create.component';
import { ProcedureOrderService } from '../../../services/ProcedureOrder.service';
import { Router } from '@angular/router';

describe('CreateProcedureOrderComponent', () => {
  let component: CreateProcedureOrderComponent;
  let fixture: ComponentFixture<CreateProcedureOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProcedureOrderComponent
      ],
      providers: [
        ProcedureOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProcedureOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});