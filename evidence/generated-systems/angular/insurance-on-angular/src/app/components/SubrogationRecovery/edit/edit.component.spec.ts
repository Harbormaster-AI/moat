
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditSubrogationRecoveryComponent } from './edit.component';
import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';

describe('EditSubrogationRecoveryComponent', () => {
  let component: EditSubrogationRecoveryComponent;
  let fixture: ComponentFixture<EditSubrogationRecoveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditSubrogationRecoveryComponent
      ],
      providers: [
        SubrogationRecoveryService,
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

    fixture = TestBed.createComponent(EditSubrogationRecoveryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});