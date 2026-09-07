
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateUnderwritingDecisionComponent } from './create.component';
import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';
import { Router } from '@angular/router';

describe('CreateUnderwritingDecisionComponent', () => {
  let component: CreateUnderwritingDecisionComponent;
  let fixture: ComponentFixture<CreateUnderwritingDecisionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateUnderwritingDecisionComponent
      ],
      providers: [
        UnderwritingDecisionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateUnderwritingDecisionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});