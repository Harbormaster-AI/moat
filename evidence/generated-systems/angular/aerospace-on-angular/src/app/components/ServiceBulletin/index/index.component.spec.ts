
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexServiceBulletinComponent } from './index.component';
import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';

describe('IndexServiceBulletinComponent', () => {
  let component: IndexServiceBulletinComponent;
  let fixture: ComponentFixture<IndexServiceBulletinComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexServiceBulletinComponent
      ],
      providers: [
        ServiceBulletinService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexServiceBulletinComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});