
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditDeviceCriterionComponent } from './edit.component';
import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';

describe('EditDeviceCriterionComponent', () => {
  let component: EditDeviceCriterionComponent;
  let fixture: ComponentFixture<EditDeviceCriterionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditDeviceCriterionComponent
      ],
      providers: [
        DeviceCriterionService,
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

    fixture = TestBed.createComponent(EditDeviceCriterionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});