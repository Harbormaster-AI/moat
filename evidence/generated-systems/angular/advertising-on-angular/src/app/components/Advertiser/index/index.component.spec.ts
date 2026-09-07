
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAdvertiserComponent } from './index.component';
import { AdvertiserService } from '../../../services/Advertiser.service';

describe('IndexAdvertiserComponent', () => {
  let component: IndexAdvertiserComponent;
  let fixture: ComponentFixture<IndexAdvertiserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAdvertiserComponent
      ],
      providers: [
        AdvertiserService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAdvertiserComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});