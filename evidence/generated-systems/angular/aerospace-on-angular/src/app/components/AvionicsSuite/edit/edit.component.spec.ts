
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditAvionicsSuiteComponent } from './edit.component';
import { AvionicsSuiteService } from '../../../services/AvionicsSuite.service';

describe('EditAvionicsSuiteComponent', () => {
  let component: EditAvionicsSuiteComponent;
  let fixture: ComponentFixture<EditAvionicsSuiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditAvionicsSuiteComponent
      ],
      providers: [
        AvionicsSuiteService,
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

    fixture = TestBed.createComponent(EditAvionicsSuiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});