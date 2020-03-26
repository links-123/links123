<template>

    <!-- header header  -->
    <div class="header is_stuck">
        <nav class="navbar top-navbar navbar-expand-md navbar-light">
            <!-- Logo -->
            <div class="navbar-header">
                <a class="navbar-brand" href="index.html">
                    <router-link to="/" class="navbar-brand invisible-xs" exact>
                        <img id="small-logo" src="../../images/logo-small.png"
                             alt="Link 1,2,3!"/>
                        <img id="full-logo" src="../../images/logo-full.png"
                             alt="Links 1,2,3!"/>
                    </router-link>
                </a>
            </div>
            <!-- End Logo -->

            <div class="navbar-collapse">
                <ul class="navbar-nav mr-auto my-lg-0">
                    <li class="nav-item">
                        <button
                                type="button"
                                class="btn btn btn-outline-success btn-sm"
                        >
                            <i class="fa fa-magic"></i>&nbsp; Add link
                        </button>
                    </li>
                </ul>
                <!-- User profile and search -->
                <ul class="navbar-nav ml-auto my-lg-0">
                    <!-- Profile -->
                    <li class="nav-item dropdown"
                        :class="{'show': isUserInfo, '': !isUserInfo }"
                        @click="handleUserInfo"
                    >
                        <a class="nav-link dropdown-toggle text-muted"
                           href="#" data-toggle="dropdown" aria-haspopup="true"
                           aria-expanded="false"><img
                                src="../../images/user.png" alt="user"
                                class="profile-pic"/></a>
                        <div
                                class="dropdown-menu dropdown-menu-right animated zoomIn"
                                :class="{'show': isUserInfo, '': !isUserInfo }"
                                @click="handleUserInfo"
                        >
                            <ul class="dropdown-user">
                                <li><a href="#"><i class="ti-settings"></i>
                                    Setting</a></li>
                                <li><a href="#"><i class="fa fa-power-off"></i>
                                    Logout</a></li>
                            </ul>
                        </div>
                    </li>
                </ul>
            </div>
        </nav>
    </div>

</template>

<script>
    export default {
        data() {
            return {
                isSearch: false,
                isOpen: false,
                clickable: true,
                isNotification: false,
                isMessage: false,
                isUserInfo: false
            }
        },
        components: {
        },
        methods: {
            toggle(e) {
                //this.isOpen = !this.isOpen
                e.target.parentElement.classList.toggle('close')
            },
            searchClick(e) {
                this.isSearch = !this.isSearch;
            },
            handleMessage(e) {
                this.isMessage = !this.isMessage;
            },
            handleNotification(e) {
                this.isNotification = !this.isNotification;
            },
            handleUserInfo(e) {
                this.isUserInfo = !this.isUserInfo;
            }
        }
    }
</script>

<style lang="scss">
    .form-inline {
        display: block !important;
    }

    .notification-show + .dropdown-menu,
    .message-show + .dropdown-menu,
    .language-show + .dropdown-menu {
        display: block;
    }

    .header .top-navbar .navbar-header .navbar-brand {
        height: 35px;
        max-width: 250px;
    }

    .navbar-brand #small-logo,
    .mini-sidebar .navbar-brand #full-logo {
        display: none;
    }

    .mini-sidebar .navbar-brand #small-logo {
        display: inline;
    }

    #small-logo, #full-logo {
        width: auto;
        height: 35px;
    }

    .header.is_stuck {
        min-height: 54px;

        z-index: 999999999;
    }

    .header.is_stuck .navbar {
        position: fixed;
        width: 100%;
        background: white;
        box-shadow: 1px 0 5px rgba(0, 0, 0, 0.1);
    }

</style> 