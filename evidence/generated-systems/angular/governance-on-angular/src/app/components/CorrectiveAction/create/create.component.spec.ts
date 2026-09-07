
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCorrectiveActionComponent } from './create.component';
import { CorrectiveActionService } from '../../../services/CorrectiveAction.service';
import { Router } from '@angular/router';

describe('CreateCorrectiveActionComponent', () => {
  let component: CreateCorrectiveActionComponent;
  let fixture: ComponentFixture<CreateCorrectiveActionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCorrectiveActionComponent
      ],
      providers: [
        CorrectiveActionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCorrectiveActionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});