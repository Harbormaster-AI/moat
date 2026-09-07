
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditLaboratoryOrderComponent } from './edit.component';
import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';

describe('EditLaboratoryOrderComponent', () => {
  let component: EditLaboratoryOrderComponent;
  let fixture: ComponentFixture<EditLaboratoryOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditLaboratoryOrderComponent
      ],
      providers: [
        LaboratoryOrderService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditLaboratoryOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});