
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditServiceBulletinComponent } from './edit.component';
import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';

describe('EditServiceBulletinComponent', () => {
  let component: EditServiceBulletinComponent;
  let fixture: ComponentFixture<EditServiceBulletinComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditServiceBulletinComponent
      ],
      providers: [
        ServiceBulletinService,
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

    fixture = TestBed.createComponent(EditServiceBulletinComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});