
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditUnderwritingDecisionComponent } from './edit.component';
import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';

describe('EditUnderwritingDecisionComponent', () => {
  let component: EditUnderwritingDecisionComponent;
  let fixture: ComponentFixture<EditUnderwritingDecisionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditUnderwritingDecisionComponent
      ],
      providers: [
        UnderwritingDecisionService,
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

    fixture = TestBed.createComponent(EditUnderwritingDecisionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});