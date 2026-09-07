
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditWorkAuthorizationComponent } from './edit.component';
import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';

describe('EditWorkAuthorizationComponent', () => {
  let component: EditWorkAuthorizationComponent;
  let fixture: ComponentFixture<EditWorkAuthorizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditWorkAuthorizationComponent
      ],
      providers: [
        WorkAuthorizationService,
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

    fixture = TestBed.createComponent(EditWorkAuthorizationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});