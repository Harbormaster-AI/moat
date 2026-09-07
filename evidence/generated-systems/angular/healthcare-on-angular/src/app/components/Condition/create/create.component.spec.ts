
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateConditionComponent } from './create.component';
import { ConditionService } from '../../../services/Condition.service';
import { Router } from '@angular/router';

describe('CreateConditionComponent', () => {
  let component: CreateConditionComponent;
  let fixture: ComponentFixture<CreateConditionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateConditionComponent
      ],
      providers: [
        ConditionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateConditionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});